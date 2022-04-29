import React, {SyntheticEvent, useState} from 'react';
import Message from '../components/Message';

import {Navigate} from "react-router-dom";

const AddPenyakit = () => {
    const [nama, setNama] = useState('');
    const [dna_seq, setDna_seq] = useState('');
    const [mes, setMes] = useState('');
    
    
    // const [redirect, setRedirect] = useState(false);
    
    

      const showFile = (e) => {
        e.preventDefault();
        const reader = new FileReader();
        reader.onload = (e) => {
          const text = e.target.result;
          console.log(text)
          setDna_seq(text);
        };
        reader.readAsText(e.target.files[0]);
      };
      



    const submit = async (e) => {
        e.preventDefault();

        const formData = new FormData();
        formData.append("nama", nama);
        formData.append("dna_seq", dna_seq);
        console.log(dna_seq);

        const response = await fetch('http://localhost:8000/api/addPenyakit', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                nama,
                dna_seq,
            })
        });

        const content = await response.json();
        console.log(content.message);

        setMes(content.message)
        
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

           

            <input type="file" className="form-control" placeholder="Sequence DNA" required accept=".txt"
                   onChange={showFile}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
        </form>

        <Message msg={mes}/>

        {/* <div>
            <h1>hello</h1>
        </div> */}
        
        
       
    </div>
    );
};

export default AddPenyakit;