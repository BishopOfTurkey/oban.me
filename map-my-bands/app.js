const request = require('request');
const parseInfo = require("infobox-parser")



function findCityString(str) {
	var url = 'http://en.wikipedia.org/w/api.php?action=query&prop=revisions&rvprop=content&format=json&rvsection=0&titles=' + str
	request(url,function (error, response, body) {
	  console.log('error:', error); // Print the error if one occurred
	  console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received
	  var page = JSON.parse(body).query.pages
	  page = page[Object.keys(page)[0]];
	  page = page.revisions[0]['*'];
	  var origin = page.match(/\| origin.*\n/)[0];
	 	origin = origin.replace(/\| origin\s*=\s/,"").replace(/[\])}[{(]/g, '');

	  console.log(origin);
	});
}




findCityString("Taylor Swift");

