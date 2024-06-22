document.addEventListener("DOMContentLoaded", () => {
    const getFieldValue = id => document.getElementById(id).value;
    const setFieldValue = (id, value) => document.getElementById(id).value = value;
    const getCredentials = () => ({
        name: getFieldValue("nameField"),
        passwd: getFieldValue("passwordField")
    });

    const handleResponse = response => {
        if (!response.ok) throw new Error('Network response was not ok');
        return response.json();
    };

    const handleError = error => console.error('Error:', error);

    const generatePassword = () => {
        fetch('/generatePassword')
            .then(handleResponse)
            .then(data => setFieldValue('passwordField', data.message))
            .catch(handleError);
    };

    const savePassword = () => {
        fetch('/save', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(getCredentials())
        })
        .then(handleResponse)
        .then(data => console.log('Password saved successfully:', data))
        .catch(handleError);
    };

    const getPasswords = () => {
        fetch('/show')
            .then(response => {
			if(!response.ok) throw new Error("Network response was not ok");
			return response.text();
		}).
		then(text => {
			setFieldValue('pwdList', text);
			})
            .catch(handleError);
    };

    document.getElementById('generatePassBtn').addEventListener('click', generatePassword);
    document.getElementById('saveBtn').addEventListener('click', savePassword);
    document.getElementById('showBtn').addEventListener('click', getPasswords);
});

