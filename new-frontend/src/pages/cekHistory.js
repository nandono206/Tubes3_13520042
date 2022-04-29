import { FormEvent, ChangeEvent, useState } from 'react';
import {
  Stack,
  FormControl,
  Input,
  Button,
  useColorModeValue,
  Heading,
  Text,
  Container,
  Flex,
  Box,
  Stat,
  StatLabel,
  StatNumber,
  SimpleGrid,
  Alert,
  AlertIcon,
} from '@chakra-ui/react';
import SimpleSidebar from '../components/SideBar.tsx';




export default function Search() {
    const [datas, setDatas] = useState('');
    const [mes, setMes] = useState([{"id": 0, "nama": "bejo"}]);


    const submit = async (e) => {
        console.log(mes);
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

    <Flex>
      
    <Flex
      minH={'100vh'}
      w = '100vw'
      align={'center'}
      justify={'center'}
      direction="column"
      bg={useColorModeValue('gray.50', 'gray.800')}
      bgColor="#7CB9E8">
      
      <Container
        maxW={'lg'}
        bg={useColorModeValue('white', 'whiteAlpha.100')}
        boxShadow={'xl'}
        rounded={'lg'}
        p={6}
        mb="50px"
        direction={'column'}>
        <Heading
          as={'h2'}
          fontSize={{ base: 'xl', sm: '2xl' }}
          textAlign={'center'}
          mb={5}>
          History Pengecekan
        </Heading>
        <form onSubmit={submit}>
            <Stack 
            direction={{ base: 'column', md: 'row' }}
            
            spacing={'12px'}
            >
                <FormControl isRequired>
                    
                    <Input required onChange={e => setDatas(e.target.value)}/>
                    
                </FormControl>
                
                <Box w={{ base: '100%', md: '40%' }}>
            <Button
            
            colorScheme={'blue'}
              w="100%"
              type="submit"
              > Search
            </Button>
          </Box>

            </Stack>

            </form>

            
        
      </Container>

      { 
                mes != null && mes[0].nama != "bejo" ?       
               
                 

                <SimpleGrid rows={{ base: 1, md: 3 }} spacing={{ base: 5, lg: 8 }}>

                {mes.map(item => (
                        

                      <Container
        maxW={'lg'}
        bg="white"
        boxShadow={'xl'}
        rounded={'lg'}
        p={6}
        direction={'column'}>
        <Heading
          as={'h2'}
          fontSize={{ base: 'xl', sm: '2xl' }}
          textAlign={'center'}
          mb={5}>
          {item.tanggal} - {item.nama_pasien} - {item.nama} - {item.status} - {(item.presentase).toFixed(2)}%
        </Heading>
        </Container>
                    ))}
        
      </SimpleGrid>
                : mes == null ?
                <Container
        maxW={'lg'}
        
        boxShadow={'xl'}
        rounded={'lg'}
        p={6}
        mb="50px"
        direction={'column'}>
          <Alert status='warning'>
            <AlertIcon />
                Tidak ada data
            </Alert> 
        </Container>
                :
                
                <h1></h1>


            }       

      
            
    </Flex>
    </Flex>
  );
}