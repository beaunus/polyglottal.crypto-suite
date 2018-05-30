const scytaleEncryptHandler = data => {
  const element = document
    .getElementById("scytale")
    .getElementsByClassName("cipher")[0]
    .getElementsByClassName("row")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};

const scytaleEncrypt = async () => {
  return await axios.get("/api/v1/scytale", {
    params: {
      plaintext: document.getElementById("scytale-encrypt-plaintext").value,
      numSides: document.getElementById("scytale-encrypt-num-sides").value
    }
  });
};

document.getElementById("scytale-encrypt-button").onclick = async () => {
  const response = await scytaleEncrypt();
  showResult(response.data, scytaleEncryptHandler);
};
