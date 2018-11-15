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
		
		let url = '/api/v1/f/upload'
		
		var ajaxData = new FormData($form.get(0));
		
		$.each( droppedFiles, function(i, file) {
			ajaxData.append( $input.attr('name'), file );
		});

		sendFile($form, ajaxData)
	});

	$("#submit").on('click', (e) => {
		e.preventDefault();
		e.stopPropagation();

		$form = $("#form")
		var ajaxData = new FormData($form.get(0));

		let url = '/api/v1/f/upload'

		$.ajax({
			url: url,
			type: 'post',
			dataType: 'json',
			cache: false,
			contentType: false,
			processData: false,
			data: ajaxData,
			success: (data) => {
				console.log(data);
				
			}
		});
		//$("#form").ajaxSubmit({url: url, type: 'POST'})
	});

	function sendFile($form, ajaxData){
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
				$('#link').html('<a href='+'/f/'+json.hash+'>u.sawol.moe/file/'+json.hash+'</a>')
				console.log(d.responseText);
				
				$form.removeClass('is-uploading');
			}
		});
	}
});