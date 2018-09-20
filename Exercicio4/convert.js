const cv = require('opencv');

cv.readImage('astley.jpg', function (err, img) {
  if (err) {
    throw err;
  }

  const width = img.width();
  const height = img.height();

  if (width < 1 || height < 1) {
    throw new Error('Image has no size');
  }

  // do some cool stuff with img
  img.convertHSVscale();
  // save img
  img.save('astleytransformed.jpg');
});
