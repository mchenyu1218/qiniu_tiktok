// import axios from 'axios'
import request from './index'

// baseURL=
// const request = axios.create({
//     baseURL:'http://localhost:8080',
//     transformResponse:[
//         function(data){
//             try{
//                 return JSON.parse(data)
//             }catch(err){
//                 return err
//             }
//         }
//     ]
// })

// export default request

// export const feedallvideo = params => {
//     return axios.get(`http://localhost:8080/api/tiktok/feedallvideo`,{ params: params });
//   };
  

export function gettag(tag){
  // console.log(tag);
  return request.get(`/feedbytag?tag=${tag}`)
}


export function getuser(userid,id){
  // console.log('userid=',id);
  return request.get(`/user/search/:UserID=${userid}?user_id=${id}`)
}