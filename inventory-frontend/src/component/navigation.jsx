import { Link } from "react-router";

function Component() {
    return (
        <nav className="navbar" role="navigation" aria-label="main navigation">
            <div id="navbarBasicExample" className="navbar-menu">
                <div className="navbar-start">
                <Link className="navbar-item" to='/'>
                    Home
                </Link>
                <Link className="navbar-item" to='/reserve'>
                    Reserve
                </Link>
                </div>
            </div>
        </nav>
    )
}

export default Component