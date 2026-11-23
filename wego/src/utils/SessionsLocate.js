export function saveInfos(tipo, token){
    let userIdentificar = [tipo, token];
    sessionStorage.setItem("valor_sessao", userIdentificar);
}

export function recoverInfos(){
    const userIdentificar = sessionStorage.getItem('valor_sessao');
    const tipo = userIdentificar[0];
    const token = userIdentificar[1];
    if (userIdentificar) {
        return tipo,token;
    }
}