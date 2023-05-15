function removePlaceholder(element: HTMLInputElement): void {
  element.setAttribute("data-placeholder", element.getAttribute("placeholder")!);
  element.removeAttribute("placeholder");
}

function restorePlaceholder(element: HTMLInputElement): void {
  element.setAttribute("placeholder", element.getAttribute("data-placeholder")!);
  element.removeAttribute("data-placeholder");
}
