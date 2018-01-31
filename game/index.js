console.log("hoge")

var point = 0
var display = document.querySelector("#display span");

document.querySelector("#cookie").addEventListener("click",function(){
  point += 1;
  console.log(point);
  display.innerHTML = point;
},false);

var submit = document.querySelector("#submit")

function sendScore(name) {
  var xhr = new XMLHttpRequest();
  xhr.onreadystatechange = function() {
    if( xhr.readyState == 4) {
      if( xhr.status == 200 || xhr.status == 304 ) {
        var data = xhr.responseText;
        console.log( 'COMPLETE! :'+data );
      } else {
        console.log( 'Failed. HttpStatus: '+xhr.statusText );
      }
    }
  };
  xhr.open( 'POST', 'http://localhost:8080/api/v1/scores' );
  xhr.setRequestHeader( 'Content-Type', 'application/json' );
  xhr.send(JSON.stringify({name: name, point: point}));
}

// ハンドラの登録.
document.querySelector("#submit").addEventListener("click",function(){
  var name = document.querySelector("#name").value;
  console.log(name);
  sendScore(name);
});
