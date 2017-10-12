import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { Link } from 'lib/nav_link'
import { graphql } from 'react-apollo'
import Notification from 'actions/notification'
import Spinner from 'components/shared/spinner'
import Page500 from 'components/shared/page500'

class Products extends Component {

  static propTypes = {
    clientsQuery: PropTypes.object.isRequired,
  }

  state = {}

  // componentWillReceiveProps(props) {
  //   let error = props.clientsQuery.error
  //   if (error) { Notification.error(error.message) }
  // }

  render() {
    // const { page } = this.props.match.params
    // let { loading, error, clients, refetch, meta } = this.props.clientsQuery

    // if (loading) {
    //   return <Spinner />
    // }

    // if (error) {
    //   return <Page500 />
    // }

    return (
      <div className="animated fadeIn">

        <div className="row">
          <div className="col-lg-12">
            <div className="card">
              <div className="card-header">
                <i className="fa fa-align-justify"></i> Products
              </div>
              <div className="card-block">
                <table className="table text-center">
                  <thead>
                    <tr>
                      <th className="text-center">Id</th>
                      <th className="text-center">Name</th>
                    </tr>
                  </thead>
                  <tbody>


                  </tbody>
                </table>

              </div>
            </div>
          </div>
        </div>

      </div>
    )
  }
}

// export default graphql(clientsQuery, {
//   name: "clientsQuery",
// })(Products)
export default Products

// { clients.map( (object, index) =>
//   <ClientView
//     key={index}
//     object={object}
//     refetch={refetch}
//   />
// )}
