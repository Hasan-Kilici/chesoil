function getCookie(name) {
    var value = "; " + document.cookie;
    var parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
  }
  
  function setCookie(name, value, days) {
    var expires = "";
    if (days) {
      var date = new Date();
      date.setTime(date.getTime() + (days*24*60*60*1000));
      expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + value + expires + "; path=/";
  }

function addToCart(productId, product) {
  var cartCookie = getCookie("cart");
  var cart = cartCookie ? JSON.parse(cartCookie) : [];
  cart.push(product);
  setCookie("cart", JSON.stringify(cart), 7);
}
function removeFromCart(productId) {
  var cartCookie = getCookie("cart");
  var cart = cartCookie ? JSON.parse(cartCookie) : [];
  var updatedCart = cart.filter(function(item) {
    return item.productid !== productId;
  });
  setCookie("cart", JSON.stringify(updatedCart), 7);
  listMyCart();
}
  function listMyCart(){
    let totalprice = 0;
    let list;
    if(getCookie("cart") != undefined){
      list = JSON.parse(getCookie("cart"));
    } else {
      list = false;
    }
    let cartBox = document.getElementById("fotolar");
    for(let i = 0;i < list.length;i++){
      cartBox.innerHTML += `<a href="/info/${list[i].id}" class="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow md:flex-row md:max-w-xl hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700">
<img class="w-full rounded-t-lg h-[96%] md:h-auto md:w-48 md:rounded-none md:rounded-l-lg" style="height:100%" src="${list[i].banner}" alt="">
<div class="flex flex-col justify-between p-4 leading-normal">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">${list[i].name}</h5>
    <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">${list[i].description}</p>
</div>
</a>`;
    }
    if(list == false){
      cartBox.innerHTML = `
  boş
      `;
      document.getElementById("priceCard").style.display = "none";
    }
  }
