import { useState, useEffect } from 'react'
import Navigation from "../component/navigation"
import Header from "../component/header"
import { BASE_API_URL } from "../utils/constant"

function Home() {
  const [inventoryStatus, setInventoryStatus] = useState({
    "item_id": 0,
    "total_stock": 0,
    "reserved_stock": 0,
    "available_stock": 0
  })

  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [currentItemId, setCurrentItemId] = useState(1)

  useEffect(() => {
    async function fetchUsers(itemId) {
      try {
        setLoading(true);
        const response = await fetch(`${BASE_API_URL}/api/v1/inventory/stock?item_id=${itemId}`);

        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        setInventoryStatus(data)
      } catch (err) {
        if (err.name !== 'AbortError') {
          setError(err.message);
        }
      } finally {
        setLoading(false);
      }
    }

    fetchUsers(currentItemId)
  }, [])

  return (
    <div className='container'>
      <Header />
      <Navigation />

      <div className='fixed-grid'>
        <div className='grid'>
          <div className='grid-cell'>
            <div className='card'>
              <div className='card-header'>
                <p className='card-header-title has-text-weight-bold'>Item ID</p>
              </div>
              <div className='card-content'>
                <p>{inventoryStatus.item_id}</p>
              </div>
            </div>
          </div>

          <div className='grid-cell'>
            <div className='card'>
              <div className='card-header'>
                <p className='card-header-title has-text-weight-bold'>Total Stock</p>
              </div>
              <div className='card-content'>
                <p>{inventoryStatus.total_stock}</p>
              </div>
            </div>
          </div>

          <div className='grid-cell'>
            <div className='card'>
              <div className='card-header'>
                <p className='card-header-title has-text-weight-bold'>Reserved Stock</p>
              </div>
              <div className='card-content'>
                <p>{inventoryStatus.reserved_stock}</p>
              </div>
            </div>
          </div>

          <div className='grid-cell'>
            <div className='card'>
              <div className='card-header'>
                <p className='card-header-title has-text-weight-bold'>Available Stock</p>
              </div>
              <div className='card-content'>
                <p>{inventoryStatus.available_stock}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home
