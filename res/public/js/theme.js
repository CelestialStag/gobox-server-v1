$("document").ready(function(){
	document.cookie = "username=John Doe";
	
	let dark = 'https://cdn.jsdelivr.net/npm/ashleycss@4.1.51/dist/themes/ashleycss-dark.min.css';
	let light = 'https://cdn.jsdelivr.net/npm/ashleycss@4.1.2/dist/ashleycss-sakura.min.css';

	let theme = Cookies.get('theme');

	if(theme ==  'dark')
	{
		$('#style').attr('href', dark);
		$('#theme').html("lights on!")
	}
	else
	{
		$('#style').attr('href', light);
		$('#theme').html("lights off!")
	}

	$('#theme').on('click', (e) => {
		e.target.value = 'qq';
		let theme = Cookies.get('theme');

		if(theme ==  'dark')
		{
			Cookies.set('theme', 'light');
			$('#style').attr('href', light);
			$('#theme').html("lights off!")
		}
		else
		{
			Cookies.set('theme', 'dark');
			$('#style').attr('href', dark);
			$('#theme').html("lights on!")
		}
	});
});