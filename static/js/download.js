$("document").ready(function(){

	$.ajax({
		url: '/api/v1/f/info/'+$("#id").attr('value'),
		type: 'GET',
		cache: false,
		contentType: false,
		processData: false,
		complete: function(d) {
			json = JSON.parse(d.responseText)
			console.log(json)
			$('#name').html(json.name)
			$('#size').html(humanFileSize(json.size, true))
			$('#expires').html(json.expires)
			$('#uploaded').html(json.uploaded)
			$('#type').html(json.type)
			$('#download').attr('value', json.url)

			if(json.type.includes("jpeg") || json.type.includes("png") || json.type.includes("gif") || json.type.includes("webp")){
				// http://localhost:4040/api/v1/f/download/1015869f/PLQ2uBL.jpg example
				$('#preview').attr('src', 'http://u.sawol.moe/api/v1/f/download/'+json.url)
			}
		}
	});

	function humanFileSize(bytes, si) {
		var thresh = si ? 1000 : 1024;
		if(Math.abs(bytes) < thresh) {
			return bytes + ' B';
		}
		var units = si
			? ['kB','MB','GB','TB','PB','EB','ZB','YB']
			: ['KiB','MiB','GiB','TiB','PiB','EiB','ZiB','YiB'];
		var u = -1;
		do {
			bytes /= thresh;
			++u;
		} while(Math.abs(bytes) >= thresh && u < units.length - 1);
		return bytes.toFixed(1)+' '+units[u];
	}

	$("#download").on('click', (e) => {
		e.preventDefault();
		let url = '/api/v1/f/download/'+e.target.value;
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