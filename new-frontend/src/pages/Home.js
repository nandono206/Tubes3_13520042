import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';
import Message from '../components/Message';

const Home = () => {
    const [datas, setDatas] = useState('');
    const [mes, setMes] = useState([{"id": 0, "nama": "bejo"}]);
    const [searched, setSearched] = useState(false);
    const [index, setIndex] = useState(0)
    
    
      
    

    const submit = async (e) => {
        e.preventDefault();

        const response = await fetch('http://localhost:8000/api/history', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                datas
            })
        });

        const content = await response.json();
        console.log(content);
        setMes(content)  
    }
    

    

    return (
        <div>
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Tes Penyakit</h1>

                <input className="form-control" placeholder="Masukkan data" required
                    onChange={e => setDatas(e.target.value)}
                />

                

                <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
            </form>

           
            {
                
                mes != null && mes[0].nama != "bejo" ? 
                
                <ul>
                    {mes.map(item => (
                        <li key={item.id}>
                          {item.tanggal} - {item.nama_pasien} - {item.nama} - {item.status} 
                        </li>
                    ))}
                </ul> : <h1>tes</h1>
            }
            
            
        </div>
    );
};

export default Home;