import "./topbar.css"
import { Link } from 'react-router-dom';
import {Salvo} from "../../utils/SessionsLocate"

conteudoEntragador = (
    <>
        <div className="link-page-topbar"><Link to="/">Home</Link></div>
        <div className="link-page-topbar"><Link to="/encomenda">Encomenda</Link></div>
        <div className="link-page-topbar"><Link to="/veiculos">Veiculos</Link></div>
        <div className="link-page-topbar"><Link to="/locais">Locais</Link></div>
        <div className="link-page-topbar"><Link to="/mensagem">Mensagem</Link></div>
    </>);

conteudoCliente = (
    <>
        <div className="link-page-topbar"><Link to="/">Home</Link></div>
        <div className="link-page-topbar"><Link to="/encomenda">Encomenda</Link></div>
        <div className="link-page-topbar"><Link to="/locais">Locais</Link></div>
        <div className="link-page-topbar"><Link to="/mensagem">Mensagem</Link></div>
    </>);              

conteudoFuncionario = (
    <>
        <div className="link-page-topbar"><Link to="/">Home</Link></div>
        <div className="link-page-topbar"><Link to="/encomenda">Encomenda</Link></div>
        <div className="link-page-topbar"><Link to="/veiculos">Veiculos</Link></div>
        <div className="link-page-topbar"><Link to="/locais">Locais</Link></div>
        <div className="link-page-topbar"><Link to="/mensagem">Mensagem</Link></div>
    </>);

if (Salvo[0] == 0){
    conteudoTipo = conteudoEntragador
} else if (Salvo[1] == 1){
    conteudoTipo = conteudoFuncionario
} else{
    conteudoTipo = conteudoCliente
}

export function TopBar(){
    return(
        <>
            <div className="conteiner-topbar">
                {conteudoTipo}
            </div>
        </>
    );
}