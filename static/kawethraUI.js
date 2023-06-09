let previousStyles = [];
let hoveredElements = [];

let objects = document.querySelectorAll("*");
let objLength = objects.length;

for (let i = 0; i < objLength; i++) {
  setClass(objects[i]);
}

function setClass(object) {
  let UI = object.classList;
  if (UI.length != 0) {
    for (let a = 0; a < UI.length; a++) {
      let Param1 = UI[a]?.split("[")[1]?.split("]")[0];
      let Param2 = Param1?.split("-")[0];
      let Param3 = UI[a]?.split("[")[2]?.split("]")[0];
      let Param4 = UI[a]?.split("[")[3]?.split("]")[0];
      let ClassValue = UI[a]?.split("[")[1]?.split("]")[0];
      let Class = UI[a]?.split("[");
      object.style[Class[0]] = ClassValue;
      if (UI[a].includes("shadow")) {
        let shadowStyle;
        switch (Param2) {
          case "sm":
            shadowStyle = "0 1px 2px 0";
            break;
          case "md":
            shadowStyle = "0 4px 6px -1px";
            break;
          case "lg":
            shadowStyle = "0 10px 15px -3px";
            break;
          case "xl":
            shadowStyle = "0 20px 25px -5px";
            break;
          default:
            shadowStyle = `${Param2} ${Param2} ${Param2} ${Param2}`;
        }
        object.style.boxShadow = `${shadowStyle} ${Param3}`;
      }
      if (UI[a].includes("border")) {
        let borderStyle;
        switch (Param2) {
          case "sm":
            borderStyle = "1px";
            break;
          case "md":
            borderStyle = "2px";
            break;
          case "lg":
            borderStyle = "4px";
            break;
          case "xl":
            borderStyle = "6px";
            break;
          default:
            borderStyle = Param2;
        }
        if (Param4) {
          object.style.border = `${borderStyle} ${Param3} ${Param4}`;
        } else {
          object.style.border = `${borderStyle} ${Param3} solid`;
        }
      } else if (UI[a].includes("hover-")) {
        let hoverProp = UI[a].split("-")[1];
        let hoverStyle = hoverProp.split("-")[2];
        let propHover = object.getAttribute("hover");
        object.setAttribute("hover", `${propHover} ${UI[a]}`);
        let classList = object.className.split(" ");
        for (let i = 0; i < classList.length; i++) {
          if (classList[i].includes(hoverProp)) {
            let Value = classList[i].match(/\[(.*?)\]/)[1];
            previousStyles.push({
              att: UI[a],
              property: hoverProp,
              style: Value,
              name: UI[a],
            });
            break;
          }
        }
      }
    }
  }
}

document.addEventListener("mouseover", (e) => {
  let element = e.srcElement;
  let hovers = element.getAttribute("hover")?.split(" ");
  if (hovers?.length < 1) {
  } else {
    for (let i = 0; i < hovers?.length; i++) {
      let hoverProp = hovers[i]?.split("-")[1];
      let hoverStyle = hovers[i]?.split("-")[2];
      element.style[hoverProp] = hoverStyle;
    }
  }
});

document.addEventListener("mouseout", (e) => {
  let element = e.srcElement;
  let hover = element.getAttribute("hover");
  let hoverText = hover;
  let hovers = hoverText?.split(" ");
  if (hovers?.length < 1) {
  } else {
    for (let i = 0; i < hovers?.length; i++) {
      let hoverProp = hovers[i]?.split("-")[1];
      let hoverStyle = hovers[i]?.split("-")[2];
      for (let a = 0; a < previousStyles?.length; a++) {
        console.log(previousStyles[a].style);
        if (previousStyles[a].name == hovers[i]) {
          console.log(previousStyles[a].property)
          console.log(previousStyles[a].style)
          element.style[previousStyles[a].property] = previousStyles[a].style;
        }
      }
    }
  }
});

window.onhashchange = ()=>{
    let objectts = document.querySelectorAll("*");
    let objtLength = objects.length;
    console.log(objtLength)
    for (let i = 0; i < objtLength; i++) {
        setClass(objects[i]);
    }
}

window.onload = ()=>{
    let objectts = document.querySelectorAll("*");
    let objtLength = objects.length;
    console.log(objtLength)
    for (let i = 0; i < objtLength; i++) {
        setClass(objects[i]);
    }   
}

window.onDOMContentLoaded = ()=>{
    let objectts = document.querySelectorAll("*");
    let objtLength = objects.length;
    console.log(objtLength)
    for (let i = 0; i < objtLength; i++) {
        setClass(objects[i]);
    }   
}