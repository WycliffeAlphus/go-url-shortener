document.getElementById('shortenForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const url = document.getElementById('urlInput').value.trim();

    if (!url) {
        alert("URL cannot be empty");
        return;
    }

    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ url }),
        });

        console.log("Response Status:", response.status);

        const text = await response.text(); // Read response as text
        console.log("Raw Response:", text);

        let data;
        try {
            data = JSON.parse(text); // Try parsing JSON
        } catch (error) {
            throw new Error("Invalid JSON response from server.");
        }

        if (data.error) {
            throw new Error(data.error);
        }

        document.getElementById('result').innerHTML = `
            <p>Short URL: <a href="${data.short_url}">${data.short_url}</a></p>
        `;
    } catch (error) {
        document.getElementById('result').innerHTML = `
            <p style="color: red;">Error: ${error.message}</p>
        `;
    }
});
