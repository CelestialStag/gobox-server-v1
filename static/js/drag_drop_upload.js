$("document").ready(function(){

	$("body").on('dragenter', (e) => {
		e.preventDefault();
        e.stopPropagation();
	});

	$("body").on('dragleave', (e) => {
		e.preventDefault();
        e.stopPropagation();
	});

	$("body").on('dragover', (e) => {
		e.preventDefault();
        e.stopPropagation();
	});

	$("html").on('drop', (e) => {
		e.preventDefault();
		e.stopPropagation();
		let dt = e.originalEvent.dataTransfer
		let files = dt.files

		console.log(files);
		
		$.each(files, (i, f) => {
			
			let url = 'http://localhost:8080/api/v1/file/upload'
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