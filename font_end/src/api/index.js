import axios from 'axios'

const request =axios.create({
    baseURL:'http://localhost:8080/api/tiktok',
    transformResponse:{
        function(res){
            try{
                return JSON.parse(res)
            }catch(err){
                return err
            }
        }
    }
}) 
export default request