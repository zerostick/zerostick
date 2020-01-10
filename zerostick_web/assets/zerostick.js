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
	$('#DeletePopupDialog').popup("open");
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
        html += '<li data-icon="plus"><a href="#" class="selectap" id="' + item.ssid + '"><h3>' + item.ssid+ '</h3><p>BSSID:'+ item.bssid + '</p></a></li>';
      });
      $('#ul_wifinetworks').html($(html));
      $('#ul_wifinetworks').trigger('create');    
      $('#ul_wifinetworks').listview('refresh');
      $('.selectap').click(function(e) {
	var ssid = $(this).attr('id');
	console.log("Item to be selected:" + ssid);
	$("#ssid").val(ssid);
	$(window).scrollTop(0);
	//$("#ssid").get(0).scrollIntoView();

      });

      
    },                                               
    error: function(msg) {              
      alert(msg.statusText);
    } 
  });  
  
}


function updateNabto() {
  $.ajax({                                                                   
    type: "GET",                                                                        
    url: "/nabto",  
    contentType: "application/json; charset=utf-8",                                                            
    dataType: "json",   
    success: function(data) {

      console.log("Got nabto config:" + JSON.stringify(data));
      $("#deviceid").val(data);
      $("#devicekey").val("");
      
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

  $("#priority").val('0');
  
  $("#addnetworkbutton").on('click',function () {
    var ssid = $("#ssid").val();
    var password = $("#password").val();
    var priority = $("#priority").val();
    if(ssid == "") {
      $('#EmptySSIDPopupDialog h3').html("Cannot add mpty SSID");
      $('#EmptySSIDPopupDialog').popup("open");
      console.log("Cannot add empty ssid");
      return;
    }
    if(priority == "") {
      $('#EmptySSIDPopupDialog h3').html("Cannot add empty Priority");
      $('#EmptySSIDPopupDialog').popup("open");
      console.log("Cannot add empty priority");
      return;
    }
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

  

  $("#updatenabtobutton").on('click',function () {
    var deviceid = $("#deviceid").val();
    var devicekey = $("#devicekey").val();

    var postdata =  {"deviceid": deviceid, "devicekey": devicekey};

    console.log("JSON: "+JSON.stringify(postdata));

    $.ajax({                                                                   
      type: 'POST',
      url: "/nabto",  
      contentType: 'application/json',
      data: JSON.stringify(postdata),
      dataType: 'json',   
      success: function(data) {
	updateNabto();
	$(".configuration-tab").hide();
	$("#nabto-tab").show();
      },                                               
      error: function(msg) {              
	alert(msg.statusText + "postdata:"+JSON.stringify(postdata));
      }
    });
    
  });

  $("#ClearACLButton").on('click',function () {
    
    $.ajax({                                                                   
      type: 'DELETE',
      url: "/nabto/delete_acl",  
      success: function(data) {
	console.log("called clear");
      },                                               
      error: function(msg) {              
	console.log("called clear.. with error");
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
    $("#priority").val('0');
    
    $(".configuration-tab").hide();
    $("#wifiaddnetwork-tab").show();
  });

  $("#nabto-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#nabto-tab").show();
    updateNabto();
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

