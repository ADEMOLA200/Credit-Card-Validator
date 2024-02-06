function validateCreditCard() {
    const cardNumber = document.getElementById('cardNumber').value;
    const expMonth = document.getElementById('expMonth').value;
    const expYear = document.getElementById('expYear').value;

    const data = {
        card_number: cardNumber,
        exp_month: expMonth,
        exp_year: expYear
    };

    fetch('http://localhost:7001', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => response.json())
    .then(result => {
        document.getElementById('result').innerHTML = result.valid ? 'Card is valid' : 'Card is invalid';
    })
    .catch(error => {
        console.error('Error:', error);
        document.getElementById('result').innerHTML = 'Error validating credit card';
    });
}

// Disable right-click
document.addEventListener('contextmenu', function (e) {
    e.preventDefault();
});

// Additional event listener for form submission using a button
document.getElementById('validateButton').addEventListener('click', function (e) {
    e.preventDefault(); // Prevent the default form submission
    validateCreditCard(); // Call the credit card validation function
});
