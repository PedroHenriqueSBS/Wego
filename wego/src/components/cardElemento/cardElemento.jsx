import "./topbar.css";
import { Link } from 'react-router';
import {recoverInfos} from "../../utils/SessionsLocate";

let tipo = recoverInfos() ;

if (tipo == 0){
    conteudoTipo = conteudoEntragador;
} else if (tipo == 1){
    conteudoTipo = conteudoFuncionario;
} else{
    conteudoTipo = conteudoCliente;
}

export function cardElemento(){
    return(
        <>
            <div className="conteiner-card">
                
            </div>
        </>
    );
}