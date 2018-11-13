$("document").ready(function(){

	$("#download").on('click', (e) => {
		e.preventDefault();
		let url = '/api/v1/file/download/'+e.target.value;
		downloadURL(url)
	});

	function downloadURL(url) {
		var hiddenIFrameID = 'hiddenDownloader',
			iframe = document.getElementById(hiddenIFrameID);
		if (iframe === null) {
			iframe = document.createElement('iframe');
			iframe.id = hiddenIFrameID;
			iframe.style.display = 'none';
			document.body.appendChild(iframe);
		}
		iframe.src = url;
	};

});