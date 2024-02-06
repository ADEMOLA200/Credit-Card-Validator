function validateCreditCard() {
    const cardNumber = document.getElementById('cardNumber').value;
    const expMonth = document.getElementById('expMonth').value;
    const expYear = document.getElementById('expYear').value;

    const data = {
        card_number: cardNumber,
        exp_month: expMonth,
        exp_year: expYear
    };

    fetch('https://ademola-creditcard-validator.netlify.app/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(result => {
        // Check if response contains valid JSON data
        if (result && result.valid !== undefined) {
            document.getElementById('result').innerHTML = result.valid ? 'Card is valid' : 'Card is invalid';
        } else {
            throw new Error('Invalid JSON data received');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        document.getElementById('result').innerHTML = 'Error validating credit card';
    });
}
