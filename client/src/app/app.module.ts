import { NgModule } from "@angular/core";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { LoginComponent } from "./login/login.component";
import { NotFoundComponent } from "./not-found/not-found.component";

@NgModule({
	declarations: [AppComponent, LoginComponent, NotFoundComponent],
	imports: [AppRoutingModule],
	bootstrap: [AppComponent],
})
export class AppModule {}
