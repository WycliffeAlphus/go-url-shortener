document.getElementById('shortenForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const url = document.getElementById('urlInput').value;
    
    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url }),
        });
        
        const data = await response.json();
        document.getElementById('result').innerHTML = `
            <p>Short URL: <a href="${data.short_url}">${data.short_url}</a></p>
        `;
    } catch (error) {
        document.getElementById('result').innerHTML = `
            <p style="color: red;">Error: ${error.message}</p>
        `;
    }
});