
function generatePassword() {
		fetch('/generatePassword')
				.then(response => response.json())
				.then((data) => {
						document.getElementById('passwordField').value = data.message;
				}).catch(error => console.error('Error:', error));
}
