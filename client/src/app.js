import React from 'react'
import { Provider } from 'react-redux'
import { ApolloProvider } from 'react-apollo'
import { history, configureStore } from 'store'
import { Route, Redirect, Switch } from 'react-router-dom'
import { ConnectedRouter } from 'react-router-redux'
import { configureClient } from 'lib/apollo_client'
// import { loadConfig } from "actions/auth"

// shared
import Layout from 'components/shared/layout'
import Page404 from 'components/shared/page404'

import Dashboard from 'components/shared/dashboard'
import Products from 'components/products/list'

export default (onUpdate) => {

  const store = configureStore()
  const client = configureClient()

  return (
    <Provider store={store}>
      <ConnectedRouter history={history}>
        <ApolloProvider store={store} client={client}>

          <Switch>

            <Layout exact path="/dashboard" name="Dashboard" component={Dashboard}/>
            <Layout exact path="/products" name="Product" component={Products}/>

            <Redirect exact from="/" to="/dashboard"/>
            <Layout path="*" component={Page404}/>
          </Switch>

        </ApolloProvider>
      </ConnectedRouter>
    </Provider>
  )
}
