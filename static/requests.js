document.addEventListener('DOMContentLoaded', function() {
  const form = document.querySelector('form');
  form.addEventListener('submit', handleFormSubmit);

  function handleFormSubmit(event) {
    event.preventDefault();

    const urlInput = document.querySelector('input[name="url"]');
    const userIdInput = document.querySelector('input[name="userId"]');

    const url = urlInput.value;
    const userId = userIdInput.value;

    const requestBody = JSON.stringify({
      long_url: url,
      user_id: userId
    });

    fetch('/create-short-url', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: requestBody
    })
      .then(response => response.json())
      .then(data => {
        const shortUrlElement = document.getElementById('short-url');
        const newDiv = document.createElement('div');
        const linkElement = document.createElement('a');
        linkElement.href = data.short_url;
        linkElement.textContent = data.short_url;
        newDiv.appendChild(linkElement);
        shortUrlElement.appendChild(newDiv);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }
});
