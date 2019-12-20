$("document").ready(function(){

	pause = false;
	fileList = []

	$("screen").on('dragenter', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(e.target).addClass("active");
	});

	$("screen").on('dragover', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(e.target).addClass("active");
	});

	$("screen").on('dragleave', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(e.target).removeClass("active");
	});

	$("screen").on('drop', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(e.target).addClass("active");
	});

	
	$("body").on('dragenter', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(".screen").addClass("active");
	});
	
	$("body").on('dragover', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(".screen").addClass("active");
	});
	
	$("body").on('dragleave', (e) => {
		e.preventDefault();
		e.stopPropagation();

		$(e.target).removeClass("active");
	});

	$("body").on('drop', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		$(".screen").removeClass("active");
		
		$form = $("#form");
		$input = $("#form input[type='file'");
		let droppedFiles = e.originalEvent.dataTransfer.files;
		
		$.each( droppedFiles, function(i, file) {
			$('#progress-copy').first().find('label').html('Upload Progress: ' + file.name);
			
			var ajaxData = new FormData($form.get(0));
			ajaxData.append( $input.attr('name'), file );

			sendFile($form, ajaxData, file);
		});


	});

	$("#submit").on('click', (e) => {
		e.preventDefault();
		e.stopPropagation();

		var $form = $("#form")
		var ajaxData = new FormData($form.get(0));
		var name = document.getElementById('file-input');
		var file = name.files.item(0);

		$('#progress-copy').first().find('label').html('Upload Progress: ' + file.name);

		
		var name = document.getElementById('file-input');
		var file = name.files.item(0);
		$("#file-label").html("no file loaded");
		
		sendFile($form, ajaxData, file);
		
		document.getElementById('file-input').value = null;
	});

	$("#file-input").on('change', (e) => {
		var name = document.getElementById('file-input');
		var file = name.files.item(0);
		if(file != null) {
			$("#file-label").html("loaded: "+file.name);
		} else {
			$("#file-label").html("no file loaded");
		}
	});

	/*
	$("#link-c button").on('click', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		let element = document.getElementById("link")
		element.select();
		document.execCommand("copy");
	});

	$("#link").on('click', (e) => {
		e.preventDefault();
		e.stopPropagation();
		
		let element = document.getElementById("link")
		element.select();
		document.execCommand("copy");
	});
	*/

	function sendFile($form, ajaxData, file){
		var list = $('#progress-list');
		var copy = $('#progress-copy').first().clone().removeClass('display-hidden');
		var link = $('#link').first().clone().removeClass('display-hidden');

		$.ajax({
			url: $form.attr('action'),
			type: $form.attr('method'),
			data: ajaxData,
			dataType: 'json',
			cache: false,
			contentType: false,
			processData: false,
			xhr: () => {

				var xhr = $.ajaxSettings.xhr();
				
				link.attr('value', file.name);

				list.append(copy);
				copy.find('label').append(link);
				
				xhr.upload.onprogress = function (e) {

					if (e.lengthComputable) {
						copy.find('progress').attr('value', (e.loaded / e.total) * 100);
					}
				};

				return xhr;
			},
			complete: function(d) {
				json = JSON.parse(d.responseText)

				copy.find('progress').removeClass('progress-primary');
				copy.find('progress').addClass('progress-success');

				link.attr('value', window.location.href+"file/"+json.hash);
				link.on('click', (e) => {
					e.preventDefault();
					e.stopPropagation();
					
					e.target.select();
					document.execCommand("copy");
				});

				//console.log(d.responseText);
			},
			error: (x, s, e) => {

				copy.find('progress').removeClass('progress-primary');
				copy.find('progress').addClass('progress-error');

				link.attr('value', s + ":" + e);
			}
		});
	}
});