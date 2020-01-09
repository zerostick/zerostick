/**
 * Some function that will be called from multiple elements
 */
function updateKnownWifiList() {
  $.ajax({                                                                   
    type: "GET",                                                                        
    url: "/wifi",  
    contentType: "application/json; charset=utf-8",                                                            
    dataType: "json",   
    success: function(data) {
      var html ='';
      console.log("Got known wifilist:" + JSON.stringify(data));
      $.each(data, function(index, item) {
        html += '<li><a href="#"><h3>' + item.ssid+ '</h3><p>Priority:'+ item.priority + '</p></a><a href="#" class="deleteap" id="' + item.ssid + '" data-icon="delete">Delete</a></li>';
      });
      $('#ul_knownwifinetworks').html($(html));
      $('#ul_knownwifinetworks').trigger('create');    
      $('#ul_knownwifinetworks').listview('refresh');
      
      // Handle delete click

      $('.deleteap').click(function(e) {
	console.log("Item to be deleted:" + $(this).attr('id'));
	$('#popupDialog').popup("open");
	$('#wifinetworkpopupid').html('SSID:'+$(this).attr('id'));
	$('#deletenetwork').data('ssid', $(this).attr('id'));
	$('#deletenetwork').off('click');
	$('#deletenetwork').on('click', function() {
	  var ssid_to_delete = $('#deletenetwork').data('ssid');
	  $.ajax({                                                                   
	    type: 'DELETE',
	    url: "/wifi/"+ssid_to_delete,
	    dataType: "text",
	    success: function(data) {
	      console.log("Deleted");
	      updateKnownWifiList();
	    },                                               
	    error: function(msg) {
	      // Delete doesn't return a proper json... so always an error
              console.log("Could no delete");
	      updateKnownWifiList();
	    }
	  });
	  
	  
	});

      });
      
    },                                               
    error: function(msg) {              
      alert(msg.statusText);
    } 
  });  
  
}


function updateWifiList() {
  $.ajax({                                                                   
    type: "GET",                                                                        
    url: "/wifilist",  
    contentType: "application/json; charset=utf-8",                                                            
    dataType: "json",   
    success: function(data) {
      var html ='';
      console.log("Got wifilist:" + JSON.stringify(data));
      $.each(data, function(index, item) {
        html += '<li data-icon="plus"><a href="#"><h3>' + item.ssid+ '</h3><p>BSSID:'+ item.bssid + '</p></a></li>';
      });
      $('#ul_wifinetworks').html($(html));
      $('#ul_wifinetworks').trigger('create');    
      $('#ul_wifinetworks').listview('refresh');
      
      
    },                                               
    error: function(msg) {              
      alert(msg.statusText);
    } 
  });  
  
}

$(document).on("pageinit", "#configuration-page", function() {

  $(document).on('click',"#configuration-gear2", function () {
    $(".configuration-tab").hide();
    $("#wifi-tab").show();
    $("#wifi-navbar").addClass("ui-btn-active");
    $("#nabto-navbar").removeClass("ui-btn-active");
    $("#zs-navbar").removeClass("ui-btn-active");
  });

  updateKnownWifiList();
  $(".configuration-tab").hide();
  $("#wifi-tab").show();
  
  $("#addnetworkbutton").on('click',function () {
    var ssid = $("#ssid").val();
    var password = $("#password").val();
    var priority = $("#priority").val();
    var postdata =  {"ssid": ssid, "password": password,"priority": 3,"use_for_sync": false};

    console.log("JSON: "+JSON.stringify(postdata));

    $.ajax({                                                                   
      type: 'POST',
      url: "/wifi",  
      contentType: 'application/json',
      data: JSON.stringify(postdata),
      dataType: 'json',   
      success: function(data) {
	updateKnownWifiList();
	$(".configuration-tab").hide();
	$("#wifi-tab").show();
      },                                               
      error: function(msg) {              
	alert(msg.statusText + "postdata:"+JSON.stringify(postdata));
      }
    });
    
  });
  
  $("#scannetworkbutton").on('click',function () {
    updateWifiList();
  });
  
  $("#addwifibutton").on('click',function () {
    updateWifiList();
    $("#ssid").val('');
    $("#password").val('');
    $("#priority").val('');
    
    $(".configuration-tab").hide();
    $("#wifiaddnetwork-tab").show();
  });

  $("#nabto-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#nabto-tab").show();
  });
  
  $("#wifi-navbar").on('click', function () {
    $(".configuration-tab").hide();
    updateKnownWifiList();
    
    $("#wifi-tab").show();
  });
  $("#zs-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#zs-tab").show();
  });

  
});

