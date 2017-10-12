import React, { Component } from 'react'
import { connect } from 'react-redux'
import { NavLink } from 'react-router-dom'

const activeRoute = (routeName, props) => {
  return props.pathname.indexOf(routeName) > -1 ? 'nav-item nav-dropdown open pointer' : 'nav-item nav-dropdown pointer'
}

const handleClick = (e) => {
  e.preventDefault();
  e.target.parentElement.classList.toggle('open')
}

class Sidebar extends Component {
  render() {
    let { login } = this.props
    const pathname = this.props.children.props.location.pathname

    return (
      <div className="sidebar">
        <nav className="sidebar-nav">
          <ul className="nav">

            <li className="nav-item">
              <NavLink
                to={'/dashboard'}
                className="nav-link"
                activeClassName="active"
              >
                <i className="icon-speedometer"></i> Dashboard <span className="badge badge-info">NEW</span>
              </NavLink>
            </li>

            <li className="nav-item">
              <NavLink
                to={'/products'}
                className="nav-link"
                activeClassName="active"
              >
                <i className="icon-puzzle"></i> Products
              </NavLink>
            </li>

          </ul>
        </nav>
      </div>
    )
  }
}

const mapStateToProps = (props) => {
  return {
    login: props.settings.login,
  }
}

export default connect(mapStateToProps)(Sidebar)
