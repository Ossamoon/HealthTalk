//
//  RootView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct RootView: View {
    
    @ObservedObject var auth = Auth.shared
    
    var body: some View {
        Group {
            if auth.token != nil {
                ContentView()
            } else {
                SignView()
            }
        }
    }
}

struct RootView_Previews: PreviewProvider {
    static var previews: some View {
        RootView()
    }
}
