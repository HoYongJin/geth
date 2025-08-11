const bip39 = require('bip39');
const HDWallet = require('ethereum-hdwallet');

function walletGenerator() {
    const mnemonic = bip39.generateMnemonic();
    console.log("Mnemonic:", mnemonic);

    const hdWallet = new HDWallet(mnemonic);


    for(let i=0; i<20; i++) {
        console.log(`${i}: 0x${hdWallet.derive(i).getAddress().toString('hex')}`);
    }
}

walletGenerator();