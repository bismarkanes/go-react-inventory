import Header from "../component/header"
import Navigation from "../component/navigation"

function Reserve() {
    return (
        <div className="container">
            <Header />
            <Navigation />

            <div className='card'>
                <p className='has-text-weight-bold'>Create Reservation</p>
            </div>
        </div>
    )
}

export default Reserve