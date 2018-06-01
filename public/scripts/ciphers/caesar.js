const caesarEncryptHandler = data => {
  const element = document
    .getElementById("caesar")
    .getElementsByClassName("cipher-method")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};

const caesarEncrypt = async () => {
  return await axios.get("/api/v1/caesar", {
    params: {
      plaintext: document.getElementById("caesar-encrypt-plaintext").value,
      alphabet: document.getElementById("caesar-encrypt-alphabet").value,
      shift: document.getElementById("caesar-encrypt-shift").value
    }
  });
};

document.getElementById("caesar-encrypt-button").onclick = async () => {
  const response = await caesarEncrypt();
  showResult(response.data, caesarEncryptHandler);
};
