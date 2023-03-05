const NotaryForm = () => ({
    contract: "",
    seed: "",
    msg_to_sign: "",
    signedMessage: "",
    signedImage: "",
    isCopied: false,
    handleClick() { 
        this.isCopied = false;
        if (this.signedMessage === '') {
            results = Alpine.store('global_funcs').go.algoSign(
                this.contract,
                this.seed,
                this.msg_to_sign
            );
            if (results.error != '') {
                console.log(results);
                alert(results.error);
            } else {
                this.signedMessage = results.data.signed_message;
                this.signedImage = results.data.base64_qr_code;
            }
        } else {
            this.signedMessage = "";
            this.signedImage = "";
            this.msg_to_sign = "";
        }
    },
    isSubmitable() {
        console.log((this.signedMessage != "") ||
        (
            this.contract != "" &&
            this.seed != "" &&
            this.msg_to_sign != ""
        ));
        return (this.signedMessage != "") ||
        (
            this.contract != "" &&
            this.seed != "" &&
            this.msg_to_sign != ""
        );
    },
    handleCopy() {
        navigator.clipboard.writeText(this.signedMessage);
        this.isCopied = true;
    },
    bind: {
        ['x-html']() { return /*html*/`
        <section>
        <form >
            <input type="text" x-model="contract" placeholder="Verification Contract Address">
            <input  type="password" x-model="seed" placeholder="Notary Secret Key">
            <textarea  type="text" x-model="msg_to_sign" placeholder="Message"></textarea>
            <button 
                :disabled="$store.app.isOnline === true || isSubmitable() === false" 
                :class="signedMessage === '' ? init_button_class : submitted_button_class", 
                @click="handleClick()" x-text="signedMessage === '' ? 'SIGN' : 'RESET'"
                type="button",
                :data-tooltip = "!$store.app.isOnline  ? 'Click to sign message' : 'Must be offline'"
            >
            </button>

        </form>
        </section>
        <section>
        <template x-if="signedMessage === ''">
        <blockquote class="--family-sans" cite="Frederick Douglass (7 December 1869)">
            <p>Mankind are not held together by lies. Trust is the foundation of society. Where there is no truth, there can be no trust, and where there is no trust, there can be no society. Where there is society, there is trust, and where there is trust, there is something upon which it is supported.</p>
        </blockquote>
        </template>
        <template x-if="signedMessage != ''">
        <figure>
            <img :style="{'max-width': '480px', 'width': '100%', 'margin-left': 'auto', 'margin-right': 'auto' }" :src="signedImage" alt="QR Code of Signed Message">
            <figcaption>QR Code of Signed Message</figcaption>
			<div class="content">
				<h3 class="title">Signed Message</h3>
                <div class="inputs">
                    <input type="text" class="--rounded-left-full" x-model="signedMessage" class="addon --rounded-right-full" aria-label="copy" readonly>
                    <button @click="handleCopy()">
                        <img height="25" :src="isCopied ? $store.images.svgs.check : $store.images.svgs.copy">
                    </button
                </div>
            </div>
        </figure>
        </template>
        </section>
        `},
    },
});

export default NotaryForm;