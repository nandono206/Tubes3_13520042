import {
    Flex,
    Stack,
    Heading,
    Text,
    Input,
    Button,
    Icon,
    useColorModeValue,
    createIcon,
    FormControl,
    FormLabel,
    AlertIcon,
  AlertTitle,
  AlertDescription,
  Alert,
  } from '@chakra-ui/react'
  import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
//   import { faCo } from "@fontawesome/free-solid-svg-icons"
import SimpleSidebar from '../components/SideBar.tsx';
  
  import React, {SyntheticEvent, useState} from 'react';
  
  export default function DnaTest() {

    const [nama, setNama] = useState('');
    const [dna_seq, setDna_seq] = useState('');
    const [nama_pasien, setNama_pasien] = useState('');
    const [mes, setMes] = useState('');
    const [pers, setPers] = useState(0.0);
    const [ket, setKet] = useState('');

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

        const response = await fetch('http://localhost:8000/api/tesSakit', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                nama,
                dna_seq,
                nama_pasien
            })
        });

        const content = await response.json();
        console.log(content.message);
        console.log(content.presentase)
        console.log(content.ket)

        setMes(content.message)
        
        console.log(ket)
        setKet(content.ket)
        console.log(ket)
        setPers((content.presentase).toFixed(2))
        
    }
    
    return (
      <Flex>
        
      <Flex
        minH={'100vh'}
        align={'center'}
        justify={'center'}
        w='100vw'
        py={12}
        
        bg={useColorModeValue('gray.50', 'gray.800')}
        bgColor="#7CB9E8">
        <Stack
          boxShadow={'2xl'}
          bg={useColorModeValue('white', 'gray.700')}
          rounded={'xl'}
          p={10}
          spacing={8}
          align={'center'}>
          <FontAwesomeIcon icon="fa-solid fa-virus-covid" />
          <Stack align={'center'} spacing={2}>
            <Heading
              textTransform={'uppercase'}
              fontSize={'3xl'}
              color={useColorModeValue('gray.800', 'gray.200')}>
              Cek Penyakit
            </Heading>
            <Text fontSize={'lg'} color={'gray.500'}>
              Mencegah Lebih Baik dari pada Mengobati Slur!
            </Text>
          </Stack>
          <form onSubmit={submit}>
            <Stack spacing={4}>
                <FormControl isRequired>
                    <FormLabel>Nama Anda</FormLabel>
                    <Input required onChange={e => setNama_pasien(e.target.value)}/>
                    
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>DNA Sequence (Must be .txt)</FormLabel>
                    <Input type="file" accept=".txt"  onChange={showFile} />
              </FormControl>
              <FormControl isRequired>
                    <FormLabel>Nama Penyakit</FormLabel>
                    <Input required onChange={e => setNama(e.target.value)}/>
                    
                </FormControl>
              <Button
                type="submit"
                  bg={'blue.400'}
                  color={'white'}
                  _hover={{
                    bg: 'blue.500',
                  }}>
                  Periksa
                </Button>

            </Stack>

            </form>
            {
                
                pers > 50.0 ? 
                
                <Alert status='error'>
                    <AlertIcon />
                        {mes} - {pers}%
                </Alert>  : ket === "succsess" ?
                <Alert status='success'>
                <AlertIcon />
                    {mes} - {pers}%
                </Alert> : ket === "error" ?

            <Alert status='warning'>
            <AlertIcon />
                {mes}
            </Alert> :
            <h1></h1>

            }
        </Stack>
      </Flex>
      </Flex>
    );
  }
  
  