package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// VideoFile has the metainfomation on each mp4
type VideoFile struct {
	Id               string // Generated UUID
	Name             string
	ThumbnailFile    string
	ThumbnailRelPath string
	FullPath         string
	Event            string
	EventType        string
	EventTime        time.Time
	EventCam         string
	Size             int64
}

// CamFS holds all files
type CamFS struct {
	VideoFiles []VideoFile
}

var (
	// ShadowCamFSPath holds where to keep temporary files
	ShadowCamFSPath = filepath.Join(os.TempDir(), "ZeroStick")
	// CamStructure holds map of path and VideoFile infomation
	CamStructure CamFS //make(map[string]VideoFile)
)

func (cfs CamFS) remove(rmFile string) {
	for i := range cfs.VideoFiles {
		if cfs.VideoFiles[i].FullPath == rmFile {
			cfs.VideoFiles[i] = VideoFile{}
		}
	}
}

// FindByID will return the VideoFile with the given `id`
func (cfs CamFS) FindByID(id string) (VideoFile, error) {
	for i := range cfs.VideoFiles {
		if cfs.VideoFiles[i].Id == id {
			return cfs.VideoFiles[i], nil
		}
	}
	return VideoFile{}, fmt.Errorf("VideoFile with Id %s not found", id)
}

func (cfs CamFS) EventsSorted() map[string]map[string][]VideoFile {
	r := make(map[string]map[string][]VideoFile)
	for i := range cfs.VideoFiles {
		//r[cfs.VideoFiles[i].Event] = append(r[cfs.VideoFiles[i].Event], cfs.VideoFiles[i])
		vf := r[cfs.VideoFiles[i].EventType][cfs.VideoFiles[i].Event]
		log.Debug("VF is", append(vf, cfs.VideoFiles[i]))
		if r[cfs.VideoFiles[i].EventType] == nil {
			r[cfs.VideoFiles[i].EventType] = map[string][]VideoFile{}
		}
		r[cfs.VideoFiles[i].EventType][cfs.VideoFiles[i].Event] = append(vf, cfs.VideoFiles[i])
	}
	log.Debug(r)
	return r
}

// HandleCamEvents will update the shadow web
func HandleCamEvents(filename string) {
	f, err := os.Stat(filename) // Get os.FileInfo
	if err != nil {             // Removed file
		CamStructure.remove(filename)
	}
	//log.Println("Found new cam file:", filename)
	indexFile(viper.GetString("cam-root")+"/TeslaCam", f) // Add f to index
}

// ScanCamFS will look trough all the files on the cam filesystem
func ScanCamFS(camfspath string) {
	// Loop over the files and directories
	files, err := ioutil.ReadDir(camfspath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() { // Dig into first level directories
			ScanCamFS(filepath.Join(camfspath, f.Name())) // Call recursive
		} else {
			indexFile(camfspath, f)
		}
	}
}

func indexFile(camfspath string, f os.FileInfo) {
	log.Debugln("New file added: ", camfspath, f.Name())
	camFSDir := camfspath[len(viper.GetString("cam-root")):] // Cut cam root off
	// Create shadow FS for thumbnails
	os.MkdirAll(filepath.Join(ShadowCamFSPath, camFSDir), 0755)

	var v VideoFile
	if f.Size() < 1000 {
		log.Warn("Ignoring ", f.Name(), "(", f.Size(), " bytes)")
	} else { // File is not corrupted
		if f.Name()[len(f.Name())-4:] == ".mp4" { // If extension matches .mp4
			v.Id = guuid.New().String()
			v.Name = f.Name()
			v.Size = f.Size()
			v.FullPath = camfspath
			v.ThumbnailRelPath = filepath.Join(camFSDir, fmt.Sprintf("%s.jpg", f.Name()))
			v.ThumbnailFile = filepath.Join(ShadowCamFSPath, v.ThumbnailRelPath)
			camFile := filepath.Join(camfspath, f.Name())
			error := GenerateCoverImage(camFile, v.ThumbnailFile, 128)
			parseFileDetails(camFile, &v) // Stuff additional details into v from the path
			if error == nil {             // Add file to index, if the was no generateCoverImage error (Which means that the file is corrupted)
				CamStructure.VideoFiles = append(CamStructure.VideoFiles, v)
			}
		}
	}
}

// SteamEnableVideo will move the `mdat` part after the `moov` metadata
// so the file is ready for streaming.
func SteamEnableVideo(videoFile string, outFile string) error {
	log.Debugln("Streaming enabling", videoFile, "as", outFile)
	// ffmpeg -i INPUT_FILE -c:v copy -crf 0 -movflags +faststart OUTPUT_FILE
	out, err := getCommandOutput("ffmpeg", "-i", videoFile, "-c:v", "copy", "-crf", "0", "-movflags", "+faststart", outFile)
	log.Debug("Conversion output:", out, err)
	return err
}

// GenerateCoverImage will genarate a cover image for the video file
// Optional image_width param will specify the width of the image.
func GenerateCoverImage(videoFile string, outFile string, imageWidth ...int) error {
	imgWidth := 1280
	if len(imageWidth) > 0 {
		imgWidth = imageWidth[0]
	}
	// ffmpeg -i video_path -vframes 1 -vf scale=$2:-2 -q:v 1 ${i%$1}jpg
	if fileExists(outFile) { // Don't make images twice
		return nil
	}
	out, err := getCommandOutput("ffmpeg", "-i", videoFile, "-vframes", "1", "-vf",
		fmt.Sprintf("scale=%d:-2", imgWidth), "-q:v", "1", outFile)
	if err != nil {
		log.Debug(out, err)
	}
	return err
}
func parseFileDetails(path string, videoFile *VideoFile) {
	// Paths has <something>/TeslaCam/SavedClips/2019-08-01_17-55-02/2019-08-01_17-56-02-front.mp4 format
	timeFormat := "2006-01-02_15-04-05" // https://stackoverflow.com/a/14106561/10334686
	rgxp := regexp.MustCompile(`/TeslaCam/((?P<EventType>.*)/(?P<Event>\d{4}-\d{2}-\d{2}_\d{2}-\d{2}-\d{2})/)?(?P<ClipTime>\d{4}-\d{2}-\d{2}_\d{2}-\d{2}-\d{2})-(?P<ClipCamera>.*).mp4$`)
	matches := rgxp.FindStringSubmatch(path)
	subnames := rgxp.SubexpNames()
	if len(matches) == 0 {
		return
	}
	// Copy matches to map
	r := make(map[string]string)
	for i, v := range subnames {
		if i == 0 {
			continue
		}
		r[v] = matches[i]
	}
	videoFile.Event = r["Event"]
	videoFile.EventType = r["EventType"]
	videoFile.EventTime, _ = time.Parse(timeFormat, r["ClipTime"])
	videoFile.EventCam = r["ClipCamera"]
}

// This is a function to execute a system command and return output
func getCommandOutput(command string, arguments ...string) (string, error) {
	// args... unpacks arguments array into elements
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	// Return the output as string and let the caller decide if and err is fatal or not
	return out.String(), err
}

// fileExists checks if a file exists and is not a directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
