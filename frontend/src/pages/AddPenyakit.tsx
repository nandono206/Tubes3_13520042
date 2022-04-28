import React, {SyntheticEvent, useState} from 'react';

import {Navigate} from "react-router-dom";

const AddPenyakit = (props: { mes: string, setMes: (mes: string) => void }) => {
    const [nama, setNama] = useState('');
    const [dna_seq, setDna_seq] = useState('');
    
    
    // const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        const response = await fetch('http://localhost:8000/api/addPenyakit', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                nama,
                dna_seq
            })
        });

        const content = await response.json();
        console.log(content.message);

        props.setMes(content.message);
        
    }

    
    

    // if (redirect) {
    //     return <Navigate to="/"/>;
    // }

    return (
    <div>
        <form onSubmit={submit}>
            <h1 className="h3 mb-3 fw-normal">Masukkan nama penyakit</h1>
            <input  className="form-control" placeholder="Nama penyakit" required
                   onChange={e => setNama(e.target.value)}
            />

            <input className="form-control" placeholder="Sequence DNA" required
                   onChange={e => setDna_seq(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>

        <div>
            <h1>{props.mes}</h1>
        </div>
    </div>
    );
};

export default AddPenyakit;