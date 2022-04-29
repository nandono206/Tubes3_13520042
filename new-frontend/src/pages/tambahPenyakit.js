import {
    Flex,
    Box,
    FormControl,
    FormLabel,
    Input,
    Checkbox,
    Stack,
    Link,
    Button,
    Heading,
    Text,
    useColorModeValue,
    Alert,
  AlertIcon,
  AlertTitle,
  AlertDescription,
  } from '@chakra-ui/react';

  import React, {SyntheticEvent, useState} from 'react';
  import SimpleSidebar from '../components/SideBar.tsx';
  
  export default function AddPenyakit() {

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


    return (
      <Flex>
        
      <Flex
        minH={'100vh'}
        w = '100vw'
        align={'center'}
        justify={'center'}
        bg={useColorModeValue('gray.50', 'gray.800')}
        bgColor="#7CB9E8">
        <Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6}>
          
          <Box
            rounded={'lg'}
            bg={useColorModeValue('white', 'gray.700')}
            boxShadow={'lg'}
            p={8}>
            <Stack spacing={4}>
            <Heading fontSize={'4xl'}>Tambahkan Penyakit</Heading>
            <Text fontSize={'m'} color={'gray.600'}>
              *Sequence DNA dan Nama Penyakit Harus Unik
            </Text>
            <form onSubmit={submit}>
            <Stack spacing={4}>
                <FormControl isRequired>
                    <FormLabel>Nama Penyakit</FormLabel>
                    <Input required onChange={e => setNama(e.target.value)}/>
                    
                </FormControl>
                <FormControl isRequired>
                    <FormLabel>DNA Sequence (Must be .txt)</FormLabel>
                    <Input type="file" accept=".txt"  onChange={showFile} />
              </FormControl>
              <Button
                type="submit"
                  bg={'blue.400'}
                  color={'white'}
                  _hover={{
                    bg: 'blue.500',
                  }}>
                  Tambah
                </Button>

            </Stack>

            </form>
            

            {
                
                mes != null && mes == "Data penyakit berhasil ditambahkan" ? 
                
                <Alert status='success'>
            <AlertIcon />
                {mes}
            </Alert>: mes != "" ?
            <Alert status='error'>
            <AlertIcon />
                {mes}
            </Alert> : <h1></h1>
            }
              
              
            </Stack>
          </Box>
        </Stack>
      </Flex>
      </Flex>
    );
  }