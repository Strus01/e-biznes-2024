import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Products = () => {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                const response = await axios.get('http://127.0.0.1:8080/products');
                setProducts(response.data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        fetchProducts();
    }, []);

    return (
        <div>
            <h1>Products</h1>
            <ul>
                {products.map(product => (
                    <li key={product.id}>{product.name} - {product.price} zł</li>
                ))}
            </ul>
        </div>
    );
};

export default Products;