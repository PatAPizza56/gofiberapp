function getFruits() {
  fetch('/api/fruits')
    .then(function(response) {
      return response.text();
    })
    .then(function(fruits) {
      alert(`Here are the fruits we found: ${fruits}`);
    });
}