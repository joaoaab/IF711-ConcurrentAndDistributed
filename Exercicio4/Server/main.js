require("sharp");

var transformer = sharp().resize(200).on('info', function(info) {
    console.log("Image height is " + info.height);
});
