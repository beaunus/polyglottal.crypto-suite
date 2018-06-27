const scytaleEncryptHandler = data => {
  console.log(data);
  const element = document
    .getElementById("scytale")
    .getElementsByClassName("cipher-method")[0]
    .getElementsByClassName("result")[0];
  element.innerText = data.Ciphertext;
};

const scytaleEncrypt = async () => {
  const plaintext = document.getElementById("scytale-encrypt-plaintext").value;
  const numSides = document.getElementById("scytale-encrypt-num-sides").value;
  console.log(plaintext, numSides);
  return await axios.get("/api/v1/scytale", {
    params: {
      plaintext,
      numSides
    }
  });
};

document.getElementById("scytale-encrypt-button").onclick = async () => {
  const response = await scytaleEncrypt();
  showResult(response.data, scytaleEncryptHandler);
};

const scytaleDecryptHandler = data => {
  const element = document
    .getElementById("scytale")
    .getElementsByClassName("cipher-method")[1]
    .getElementsByClassName("result")[0];
  element.innerText = data.Plaintext;
};

const scytaleDecrypt = async () => {
  const ciphertext = document.getElementById("scytale-decrypt-ciphertext")
    .value;
  const numSides = document.getElementById("scytale-decrypt-num-sides").value;
  console.log(ciphertext, numSides);
  return await axios.get("/api/v1/scytale", {
    params: {
      ciphertext,
      numSides
    }
  });
};

document.getElementById("scytale-decrypt-button").onclick = async () => {
  const response = await scytaleDecrypt();
  showResult(response.data, scytaleDecryptHandler);
};
