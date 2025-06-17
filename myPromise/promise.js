const fs = require('fs');

class myPromise{
      
   constructor(fn){
     //lexical scoping ; 
     const afterDone = ((data)=>{this.resolve(data)})
     const notDone  =(() =>{})
       
     fn(afterDone,notDone);    
  }
  then(cb){
      this.resolve = cb; 
  }
}


function fun(resolve, reject){
       fs.readFile('a.txt','utf-8', function (err, data){
       if(err){console.log(err)}
       resolve(data)
  })
}

function callback(data){
console.log(`everything is fine we finallly made it , Promise class! ${data}` )}


let p = new myPromise(fun); 

p.then(callback);

