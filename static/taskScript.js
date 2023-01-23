window.onload = updateCheck;

function updateCheck() {
var myNodelist = document.getElementsByTagName("LI");
var i;
for (i = 0; i < myNodelist.length; i++) {
    if(myNodelist[i].getAttribute("data")=="true")
        myNodelist[i].classList.toggle('checked');
}

}

function signOutFunc()
{
    window.location.href = "/signOut";
}


// Create a "close" button and append it to each list item
var myNodelist = document.getElementsByTagName("LI");
var i;
for (i = 0; i < myNodelist.length; i++) {
  var span = document.createElement("SPAN");
  var txt = document.createTextNode("\u00D7");
  span.className = "close";
  span.appendChild(txt);
  myNodelist[i].appendChild(span);
}

// Click on a close button to hide the current list item
var close = document.getElementsByClassName("close");
var i;
for (i = 0; i < close.length; i++) {
  close[i].onclick = function() {
    var div = this.parentElement;
    var finish_now=div.getAttribute("data")=="false"?false:true;
    var detail_now=div.innerText.slice(0, -2); 
    div.remove();

    fetch('/deleteTask', {
      method: 'Delete',
      headers: {
        'Accept': 'application/json, text/plain, */*',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ID:parseInt(div.getAttribute("data-id")) ,Finished: finish_now, Detail: detail_now})
    }).catch(e => {
        console.log(e);
    });

  }
}

// Add a "checked" symbol when clicking on a list item
var list = document.querySelector('ul');
list.addEventListener('click', function(ev) {
  var finish_cur=ev.target.getAttribute("data")=="true"?true:false;
  if (ev.target.tagName === 'LI') {
    ev.target.classList.toggle('checked');
    if(ev.target.getAttribute("data")=="false")
      ev.target.setAttribute("data","true");
    else ev.target.setAttribute("data","false");
 
  fetch('/updateTask', {
    method: 'put',
    headers: {
      'Accept': 'application/json, text/plain, */*',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ID:parseInt(ev.target.getAttribute("data-id")),Finished: finish_cur, Detail:ev.target.innerText.slice(0, -2) })
  }).catch(e => {
      console.log(e);
  });
}


}, false);

// Create a new list item when clicking on the "Add" button
function newElement() {
  var li = document.createElement("li");
  var inputValue = document.getElementById("myInput").value;
  var t = document.createTextNode(inputValue);
  li.appendChild(t);
  var empty=true;
  if (inputValue === '') {
    alert("You must write something!");
  } else {
    document.getElementById("myUL").appendChild(li);
    empty=false;
  }
  document.getElementById("myInput").value = "";

  var span = document.createElement("SPAN");
  var txt = document.createTextNode("\u00D7");
  span.className = "close";
  span.appendChild(txt);
  li.appendChild(span);

  for (i = 0; i < close.length; i++) {
    close[i].onclick = function() {
      var div = this.parentElement;
      div.style.display = "none";
    }
  }
if(!empty){
  fetch('/addTask', {
    method: 'POST',
    headers: {
      'Accept': 'application/json, text/plain, */*',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({Finished: false, Detail: inputValue})
  }) .then(response => response.json())
      .then(response => li.setAttribute("data-id",JSON.parse(JSON.stringify(response)).Task.ID), li.setAttribute("data",false));


}


}



window.addEventListener( "pageshow", function ( event ) {
    var historyTraversal = event.persisted || 
                           ( typeof window.performance != "undefined" && 
                                window.performance.navigation.type === 2 );
    if ( historyTraversal ) {
      // Handle page restore.
      window.location.reload();
    }
  });

  var input = document.getElementById("myInput");
input.addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
        event.preventDefault();
        newElement();
    }
});