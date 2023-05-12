function removePlaceholder(element) {
    element.setAttribute("data-placeholder", element.getAttribute("placeholder"));
    element.removeAttribute("placeholder");
  }
  
  function restorePlaceholder(element) {
    element.setAttribute("placeholder", element.getAttribute("data-placeholder"));
    element.removeAttribute("data-placeholder");
  }
  