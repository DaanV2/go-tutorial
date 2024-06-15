function getTime() {
    fetch('/api/time')
        .then(response => response.text())
        .then(time => {
            document.getElementById('time').innerText = time;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}