$("document").ready(function(){

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
		
		$form = $("#form")
		$input = $("#form input[type='file'")
		let droppedFiles = e.originalEvent.dataTransfer.files
		
		let url = '/api/v1/file/upload'
		
		var ajaxData = new FormData($form.get(0));
		
		$.each( droppedFiles, function(i, file) {
			ajaxData.append( $input.attr('name'), file );
		});

		$.ajax({
			url: $form.attr('action'),
			type: $form.attr('method'),
			data: ajaxData,
			dataType: 'json',
			cache: false,
			contentType: false,
			processData: false,
			complete: function(d) {
				json = JSON.parse(d.responseText)
				$.each(json, (i, v) => {
				});
				$('#link').html('<a href='+'/file/'+json.hash+'>u.sawol.moe/file/'+json.hash+'</a>')
				console.log(d.responseText);
				
				$form.removeClass('is-uploading');
			}
		});
	});
});