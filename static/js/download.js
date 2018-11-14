$("document").ready(function(){

	$.ajax({
		url: '/api/v1/file/list/'+$("#title").attr('value'),
		type: 'GET',
		cache: false,
		contentType: false,
		processData: false,
		complete: function(d) {
			json = JSON.parse(d.responseText)
			$.each(json, (i, v) => {
				console.log(v)
				$('#files').append('<li>'+v+'</li>')
				$('#files li:last-child').append('<a href=\'/api/v1/file/download/'+$("#title").attr('value')+'/'+v+'\'/></a>')
				$('#files a:last-child').append('<button>Download</button>')	
			})
			
			console.log(json[1]);
		}
	});

	// $("#download").on('click', (e) => {
	// 	e.preventDefault();
	// 	let url = '/api/v1/file/download/'+e.target.value;
	// 	downloadURL(url)
	// });

	// function downloadURL(url) {
	// 	var hiddenIFrameID = 'hiddenDownloader',
	// 		iframe = document.getElementById(hiddenIFrameID);
	// 	if (iframe === null) {
	// 		iframe = document.createElement('iframe');
	// 		iframe.id = hiddenIFrameID;
	// 		iframe.style.display = 'none';
	// 		document.body.appendChild(iframe);
	// 	}
	// 	iframe.src = url;
	// };

});