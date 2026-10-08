import { useState, useEffect } from 'react'

import Header from "../component/header"
import Navigation from "../component/navigation"
import { BASE_API_URL } from "../utils/constant"

const initialReserve = {
  status: "failed",
  reservation_id: "",
  item_id: "",
  user_id: "",
  quantity: 0,
}

function Reserve() {
  const [reserve, setReserve] = useState(initialReserve)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [refetch, setRefetch] = useState(0)
  function updateRefetch() {
    let n = refetch + 1
    setRefetch(n)
  }

  const [inputUserId, setInputUserId] = useState('');
  function handleInputUserIdChange(event) {
    setInputUserId(event.target.value)
  }

  const [inputQuantity, setInputQuantity] = useState('');
  function handleInputQuantityChange(event) {
    setInputQuantity(event.target.value)
  }

  async function confirmReserveButtonClick() {
    let body = {
      reservation_id: reserve.reservation_id
    }

    async function confirmReserveById(body) {
      try {
        setLoading(true);
        const response = await fetch(
          `${BASE_API_URL}/api/v1/inventory/confirm`,
          {
            method: 'POST',
            headers: {
              'content-type': 'application/json'
            },
            body: JSON.stringify(body),
          }
        );

        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        updateRefetch()
      } catch (err) {
        if (err.name !== 'AbortError') {
          setError(err.message);
        }
      } finally {
        setLoading(false);
      }
    }

    await confirmReserveById(body)
  }

  async function createReserveButtonClick() {
    let body = {
      user_id: inputUserId,
      item_id: "1",
      quantity: parseInt(inputQuantity)
    }

    async function createReserveByItemId(body) {
      try {
        setLoading(true);
        const response = await fetch(
          `${BASE_API_URL}/api/v1/inventory/reserve`,
          {
            method: 'POST',
            headers: {
              'content-type': 'application/json'
            },
            body: JSON.stringify(body),
          }
        );

        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        setInputUserId('')
        setInputQuantity('')
        updateRefetch()
      } catch (err) {
        if (err.name !== 'AbortError') {
          setError(err.message);
        }
      } finally {
        setLoading(false);
      }
    }

    await createReserveByItemId(body)
  }

  useEffect(() => {
    async function fetchReserveByItemId(itemId) {
      try {
        setLoading(true);
        const response = await fetch(`${BASE_API_URL}/api/v1/inventory/active?item_id=${itemId}`);

        if (!response.ok) {
          setReserve(initialReserve)
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        setReserve(data)
      } catch (err) {
        if (err.name !== 'AbortError') {
          setError(err.message);
        }
      } finally {
        setLoading(false);
      }
    }

    fetchReserveByItemId(1)
  }, [refetch])

  return (
    <div className="container">
      <Header />
      <Navigation />
      <div className='card'>
        <p className='has-text-weight-bold'>Current Reservation</p>
        { reserve.status == "success" ?
          <>
            <p>Active Reserve ID is {reserve.reservation_id}</p>
            <p>Quantity is {reserve.quantity}</p>
            <p>Please hit button to confirm</p>
            <button onClick={confirmReserveButtonClick} className='button'>Confirm</button>
          </> : null }
      </div>

      <div className='card'>
        <p className='has-text-weight-bold'>Create Quantity Reservation</p>
        <input className='input' type='text' placeholder='Enter User Id' value={inputUserId} onChange={handleInputUserIdChange} />
        <input className='input' type='number' placeholder='Enter quantity' value={inputQuantity} onChange={handleInputQuantityChange}/>
        <button onClick={createReserveButtonClick} className='button'>Create</button>
      </div>
    </div>
  )
}

export default Reserve
