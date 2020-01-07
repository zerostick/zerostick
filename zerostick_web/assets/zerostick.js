

$(document).on("pageinit", "#configuration-page", function() {
  $(".configuration-tab").hide();
  $("#wifi-tab").show();

  $("#nabto-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#nabto-tab").show();
  });
  $("#wifi-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#wifi-tab").show();
  });
  $("#zs-navbar").on('click', function () {
    $(".configuration-tab").hide();
    $("#zs-tab").show();
  });

  
});

