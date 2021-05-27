import { NgModule } from "@angular/core";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { LoginComponent } from "./login/login.component";

@NgModule({
	declarations: [AppComponent, LoginComponent],
	imports: [AppRoutingModule],
	bootstrap: [AppComponent],
})
export class AppModule {}
