import React, { useState } from 'react';
import axios from 'axios';

let id_track = 0;

const Payments = () => {
    const [status, setStatus] = useState('');

    const handlePayment = async () => {
        try {
            await axios.post('http://127.0.0.1:8080/payments', {
                id: id_track,
                price: 100,
                paymentMethod: 'blik'
            });
            setStatus('Payment successful');
            id_track += 1;
        } catch (error) {
            setStatus('Payment failed');
            console.error('Error sending payment:', error);
        }
    };

    return (
        <div>
            <h1>Payments</h1>
            <button onClick={handlePayment}>Pay 100 zł</button>
            <p>Status: {status}</p>
        </div>
    );
};

export default Payments;