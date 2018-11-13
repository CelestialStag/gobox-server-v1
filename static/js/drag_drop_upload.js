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

		let dt = e.originalEvent.dataTransfer
		let files = dt.files

		console.log(files);
		
		$.each(files, (i, f) => {
			
			let url = '/api/v1/file/upload'
			let formData = new FormData()

			formData.append('uploadfile', f)
			
			console.log(formData)

			fetch(url, {
				method: 'POST',
				body: formData
			})
		})
	});
});