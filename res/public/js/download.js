$("document").ready(function(){

	$('#logo img').addClass("small-img");

	$.ajax({
		url: '/api/file/info/'+$("#id").attr('value'),
		type: 'GET',
		cache: false,
		contentType: false,
		processData: false,
		complete: function(d) {
			json = JSON.parse(d.responseText)
			console.log(json)
			$('#name').html(json.name)
			// $('#title').html("FILE: "+json.name)
			$('#size').html(humanFileSize(json.size, true))
			$('#expires').html(json.expires)
			$('#uploaded').html(json.uploaded)
			$('#type').html(json.type)
			$('#download').attr('value', json.url)

			if(json.type.includes("jpeg") || json.type.includes("png") || json.type.includes("gif") || json.type.includes("webp") || json.type.includes("jpg") || json.type.includes("svg") || json.type.includes("apng")){
				$('#preview').attr('src', '/api/file/download/'+json.url)

				// meta
				// $('title').html("GoBox: "+json.name)
				// $('#title').attr('content', "GoBox: "+json.name);
				// $('#og-title').attr('content', "GoBox: "+json.name);
				// $('#tw-title').attr('content', "GoBox: "+json.name);
	
				// $('#desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
				// $('#og-desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
				// $('#tw-desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
	
				// $('#og-img').attr("content", '/api/file/download/'+json.url);
				// $('#tw-img').attr("content", '/api/file/download/'+json.url);
				// $('#ap-img').attr("content", '/api/file/download/'+json.url);
			}

			if(json.type.includes("webm") || json.type.includes("mp4") || json.type.includes("mkv") || json.type.includes("x-matroska") || json.type.includes("ogv") || json.type.includes("ogg")){
				$('#preview-video').attr('src', '/api/file/download/'+json.url)
				$('#preview-video').addClass("active");

				// meta
				// $('title').html("GoBox: "+json.name)
				// $('#title').attr('content', "GoBox: "+json.name);
				// $('#og-title').attr('content', "GoBox: "+json.name);
				// $('#tw-title').attr('content', "GoBox: "+json.name);
	
				// $('#desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
				// $('#og-desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
				// $('#tw-desc').attr("content", "Download file: "+json.name+" at gobox.dev. Free private, high speed file hosting")
	
				// $('#og-img').attr("content", '/api/file/download/'+json.url);
				// $('#tw-img').attr("content", '/api/file/download/'+json.url);
				// $('#ap-img').attr("content", '/api/file/download/'+json.url);
			}
		},
		error: (x, s, e) => {
			
			if(x.status == 404) {
				$('#content').html(s + ': ' + e + '\n <small>file probably does not exist or was deleted in a wipe</small>');
			} else if (x.status == 500) {
				$('#content').html(s + ': ' + e + '\n <small>error reading file from disk <br> (reloading the page might help)</small>');
			} else {
				$('#content').html(s + ': ' + e);
			}
		}
	});

	function humanFileSize(bytes, si) {
		var thresh = si ? 1000 : 1024;
		if(Math.abs(bytes) < thresh) {
			return bytes + ' B';
		}
		var units = si
			? ['KB','MB','GB','TB','PB','EB','ZB','YB']
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
		let url = '/api/file/download/'+e.target.value;
		console.log(url);
		
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