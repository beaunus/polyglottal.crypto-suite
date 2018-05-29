const showResult = (data, handler) => {
  handler(data);
};
const caesarCipherEncryptHandler = data => {
  const element = document
    .getElementById("caesar-cipher")
    .getElementsByClassName("encrypt")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};
window.onload = () => {
  document.getElementById(
    "caesar-cipher-encrypt-button"
  ).onclick = async () => {
    const response = await axios.get("/caesarCipher", {
      params: {
        plaintext: document.getElementById("caesar-cipher-encrypt-plaintext")
          .value,
        shift: document.getElementById("caesar-cipher-encrypt-shift").value
      }
    });
    showResult(response.data, caesarCipherEncryptHandler);
  };
};
